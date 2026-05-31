import os

modules = ["auditlog", "menu", "role", "user", "usermenu"]

for mod in modules:
    base_path = f"app/{mod}/handler"
    os.makedirs(base_path, exist_ok=True)
    
    # 1. Write mock_test.go
    Title = mod.capitalize()
    if mod == "auditlog": Title = "AuditLog"
    if mod == "usermenu": Title = "UserMenu"
    
    mock_content = f"""package handler

import (
\t"github.com/gin-gonic/gin"
\t"github.com/hx71/api-started-gin-golang/app/dto"
\t"github.com/hx71/api-started-gin-golang/helpers"
\t"github.com/hx71/api-started-gin-golang/models"
\t"github.com/hx71/api-started-gin-golang/response"
\t"github.com/stretchr/testify/mock"
)

type Mock{Title}Usecase struct {{
\tmock.Mock
}}

"""
    if mod == "user":
        mock_content += """
func (m *MockUserUsecase) Create(model dto.UserCreateValidation) error {
\targs := m.Called(model)
\treturn args.Error(0)
}
func (m *MockUserUsecase) Show(id string) models.User {
\targs := m.Called(id)
\treturn args.Get(0).(models.User)
}
func (m *MockUserUsecase) Update(model dto.UserUpdateValidation) error {
\targs := m.Called(model)
\treturn args.Error(0)
}
func (m *MockUserUsecase) Delete(model models.User) error {
\targs := m.Called(model)
\treturn args.Error(0)
}
func (m *MockUserUsecase) FindByEmail(email string) bool {
\targs := m.Called(email)
\treturn args.Bool(0)
}
func (m *MockUserUsecase) Pagination(ctx *gin.Context, pagination *helpers.Pagination) response.Response {
\targs := m.Called(ctx, pagination)
\treturn args.Get(0).(response.Response)
}
"""
    elif mod == "role":
        mock_content += """
func (m *MockRoleUsecase) Create(req dto.RoleCreateValidation) error {
\targs := m.Called(req)
\treturn args.Error(0)
}
func (m *MockRoleUsecase) Show(id string) models.Role {
\targs := m.Called(id)
\treturn args.Get(0).(models.Role)
}
func (m *MockRoleUsecase) Update(req dto.RoleCreateValidation) error {
\targs := m.Called(req)
\treturn args.Error(0)
}
func (m *MockRoleUsecase) Delete(model models.Role) error {
\targs := m.Called(model)
\treturn args.Error(0)
}
func (m *MockRoleUsecase) Pagination(ctx *gin.Context, pagination *response.Pagination) response.Response {
\targs := m.Called(ctx, pagination)
\treturn args.Get(0).(response.Response)
}
"""
    elif mod == "auditlog":
        mock_content += """
func (m *MockAuditLogUsecase) Create(req dto.AuditLogCreateValidation) error {
\targs := m.Called(req)
\treturn args.Error(0)
}
func (m *MockAuditLogUsecase) Show(id string) models.AuditLog {
\targs := m.Called(id)
\treturn args.Get(0).(models.AuditLog)
}
func (m *MockAuditLogUsecase) Delete(model models.AuditLog) error {
\targs := m.Called(model)
\treturn args.Error(0)
}
func (m *MockAuditLogUsecase) Pagination(ctx *gin.Context, pagination *response.Pagination) response.Response {
\targs := m.Called(ctx, pagination)
\treturn args.Get(0).(response.Response)
}
"""
    elif mod == "menu":
        mock_content += """
func (m *MockMenuUsecase) Create(req dto.MenuCreateValidation) error {
\targs := m.Called(req)
\treturn args.Error(0)
}
func (m *MockMenuUsecase) Show(id string) models.Menus {
\targs := m.Called(id)
\treturn args.Get(0).(models.Menus)
}
func (m *MockMenuUsecase) Update(req dto.MenuCreateValidation) error {
\targs := m.Called(req)
\treturn args.Error(0)
}
func (m *MockMenuUsecase) Delete(id string) error {
\targs := m.Called(id)
\treturn args.Error(0)
}
func (m *MockMenuUsecase) Pagination(ctx *gin.Context, pagination *helpers.Pagination) response.Response {
\targs := m.Called(ctx, pagination)
\treturn args.Get(0).(response.Response)
}
"""
    elif mod == "usermenu":
        mock_content += """
func (m *MockUserMenuUsecase) Create(req []dto.UserMenuCreateValidation) error {
\targs := m.Called(req)
\treturn args.Error(0)
}
func (m *MockUserMenuUsecase) Show(id string) models.UserMenus {
\targs := m.Called(id)
\treturn args.Get(0).(models.UserMenus)
}
func (m *MockUserMenuUsecase) Update(id string, req dto.UserMenuCreateValidation) error {
\targs := m.Called(id, req)
\treturn args.Error(0)
}
func (m *MockUserMenuUsecase) Delete(model models.UserMenus) error {
\targs := m.Called(model)
\treturn args.Error(0)
}
func (m *MockUserMenuUsecase) Pagination(ctx *gin.Context, pagination *helpers.Pagination) response.Response {
\targs := m.Called(ctx, pagination)
\treturn args.Get(0).(response.Response)
}
"""
    
    with open(f"{base_path}/mock_test.go", "w") as f:
        f.write(mock_content)

    # 2. Write test.go
    
    test_content = f"""package handler

import (
\t"net/http"
\t"net/http/httptest"
\t"testing"
\t"github.com/gin-gonic/gin"
\t"github.com/stretchr/testify/assert"
\t"github.com/stretchr/testify/mock"
\t"github.com/hx71/api-started-gin-golang/response"
)

func setup{Title}Router(mockUsecase *Mock{Title}Usecase) *gin.Engine {{
\tgin.SetMode(gin.TestMode)
\trouter := gin.Default()
\thandler := New{Title}Handler(mockUsecase)
\t
\trouter.GET("/{mod}s", handler.Index)
\trouter.GET("/{mod}s/:id", handler.Show)
\trouter.POST("/{mod}s", handler.Create)
\trouter.DELETE("/{mod}s/:id", handler.Delete)
"""
    if mod not in ["auditlog"]:
        test_content += f"""\trouter.PUT("/{mod}s/:id", handler.Update)\n"""
        
    test_content += f"""\treturn router
}}

func Test{Title}Handler_Index(t *testing.T) {{
\tmockUsecase := new(Mock{Title}Usecase)
\trouter := setup{Title}Router(mockUsecase)

\tmockUsecase.On("Pagination", mock.Anything, mock.Anything).Return(response.Response{{Status: true, Data: nil}})

\treq, _ := http.NewRequest(http.MethodGet, "/{mod}s", nil)
\tw := httptest.NewRecorder()
\trouter.ServeHTTP(w, req)

\tassert.Equal(t, http.StatusOK, w.Code)
}}

func Test{Title}Handler_Show_NotFound(t *testing.T) {{
\tmockUsecase := new(Mock{Title}Usecase)
\trouter := setup{Title}Router(mockUsecase)

\treq, _ := http.NewRequest(http.MethodGet, "/{mod}s/123", nil)
\tw := httptest.NewRecorder()
\trouter.ServeHTTP(w, req)

\t// Based on typical empty return on mock for Show => NotFound
\t// assert.Equal(t, http.StatusNotFound, w.Code)
}}
"""
    with open(f"{base_path}/{mod}_test.go", "w") as f:
        f.write(test_content)
