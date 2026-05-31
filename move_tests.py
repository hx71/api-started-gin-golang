import os
import glob
import re

os.makedirs("test", exist_ok=True)

modules = ["auth", "auditlog", "menu", "role", "user", "usermenu"]

for mod in modules:
    # Move and refactor mock_test.go
    mock_src = f"app/{mod}/handler/mock_test.go"
    mock_dst = f"test/mock_{mod}_test.go"
    if os.path.exists(mock_src):
        with open(mock_src, "r") as f: content = f.read()
        content = content.replace("package handler", "package test")
        
        # Add import to the usecase/dto/models if missing
        imports_to_add = [
            '"github.com/hx71/api-started-gin-golang/app/dto"',
            '"github.com/hx71/api-started-gin-golang/models"',
            '"github.com/hx71/api-started-gin-golang/response"',
            '"github.com/hx71/api-started-gin-golang/helpers"',
            '"github.com/gin-gonic/gin"',
            '"github.com/golang-jwt/jwt"',
        ]
        
        for imp in imports_to_add:
            if imp not in content:
                content = content.replace("import (", f"import (\n\t{imp}")

        with open(mock_dst, "w") as f: f.write(content)
        os.remove(mock_src)
        
    # Move and refactor {mod}_test.go
    test_src = f"app/{mod}/handler/{mod}_test.go"
    test_dst = f"test/{mod}_test.go"
    if os.path.exists(test_src):
        with open(test_src, "r") as f: content = f.read()
        
        content = content.replace("package handler", "package test")
        
        handler_import = f'"github.com/hx71/api-started-gin-golang/app/{mod}/handler"'
        if handler_import not in content:
            content = content.replace("import (", f"import (\n\thandlerPackage {handler_import}")
            
        # replace NewXHandler to handlerPackage.NewXHandler
        Title = mod.capitalize()
        if mod == "auditlog": Title = "AuditLog"
        elif mod == "usermenu": Title = "UserMenu"
        
        content = content.replace(f"New{Title}Handler", f"handlerPackage.New{Title}Handler")
        
        with open(test_dst, "w") as f: f.write(content)
        os.remove(test_src)
