package v1

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/mengdong123/go-gin-blog-jianyu/models"
	"github.com/mengdong123/go-gin-blog-jianyu/pkg/e"
	"github.com/mengdong123/go-gin-blog-jianyu/pkg/setting"
	"github.com/mengdong123/go-gin-blog-jianyu/pkg/util"
	"github.com/unknwon/com"
	"net/http"
)

// 获取文章标签
func GetTags(c *gin.Context) {
	data := make(map[string]interface{})
	maps := make(map[string]interface{})

	name := c.Query("name")
	if name != "" {
		maps["name"] = name
	}

	state := -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	code := e.SUCCESS
	// 进行查询
	data["total"] = models.GetTagTotal(maps)
	data["lists"] = models.GetTags(util.GetPage(c), setting.AppSetting.PageSize, maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// 根据id获取标签
func GetTag(c *gin.Context) {

}

// 删除文章标签
func DeleteTag(c *gin.Context) {
	// 1. 获取id
	id := com.StrTo(c.Param("id")).MustInt()

	// 验证id的值
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID最小值为1")

	code := e.INVALID_PARAMS
	// 2. 作删除操作
	if !valid.HasErrors() {
		models.DeleteTagById(id)
		code = e.SUCCESS
	} else {
		code = e.ERROR_NOT_EXIST_TAG
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]interface{}),
	})

}

// 修改文章标签
func EditTag(c *gin.Context) {
	// 1. 从context中获取id，name，modified_byd的值
	// c.query接收url中？后面的参，c.param接收url中类似id之类（/:id）的参数
	id := com.StrTo(c.Param("id")).MustInt()
	name := c.Query("name")
	modifiedBy := c.Query("modified_by")

	valida := validation.Validation{}

	// 2，获取state值并对其进行判断
	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		// https://www.cnblogs.com/kaituorensheng/p/12271691.html ：关于range的用法
		valida.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}

	// 3. 验证必要输入的值，异常的话抛错
	valida.Required(id, "id").Message("ID不能为空")
	valida.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valida.MaxSize(modifiedBy, 100, "modified_by").Message("修改人长度最长为100")
	valida.MaxSize(name, 100, "name").Message("标题长度最长为100")

	code := e.INVALID_PARAMS
	// 4. 进行更新
	if !valida.HasErrors() {
		// 判斷id的標題是否存在
		if models.ExistTagByID(id) {
			// 封装要查询的数据
			data := make(map[string]interface{})
			data["id"] = id
			if name != "" {
				data["name"] = name
			}
			if state != -1 {
				data["state"] = state
			}
			models.EditTag(id, data)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	}
	// 5. 返回
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

// 增加文章标签
func AddTag(c *gin.Context) {
	name := c.Query("name")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	createdBy := c.Query("created_by")

	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.MaxSize(createdBy, 100, "created_by").Message("创建人最长为100字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if !models.ExistTagByName(name) {
			code = e.SUCCESS
			models.AddTag(name, state, createdBy)
		} else {
			code = e.ErrorExistTag
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}
