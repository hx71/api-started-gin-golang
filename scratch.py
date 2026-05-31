import os
import glob
import re

handlers = glob.glob('app/*/handler/*.go')
for path in handlers:
    if 'test' in path: continue
    with open(path, 'r') as f:
        content = f.read()

    # standardize Data not found
    content = content.replace('"Data not found", "No data with given id"', '"data not found", "no data with given id"')
    content = content.replace('"Data not found"', '"data not found"')
    content = content.replace('"No data with given id"', '"no data with given id"')
    
    # remove commented code
    # There are commented out block in user.go like:
    # 	// if !u.Usecase.FindByEmail(req.Email) {
    # 	// 	response := response.ResponseError(config.MessageErr.FailedProcess, "duplicate email")
    # 	// 	ctx.JSON(http.StatusConflict, response)
    # 	// } else {
    # and the trailing:
    # 	// }
    
    # Wait, the instruction is also to standardize other strings if needed.
    # We will just do a regex replace for the commented out if block in user.go
    content = re.sub(r'\t// if !u\.Usecase\.FindByEmail[^\n]*\n\t// \tresponse := response\.ResponseError[^\n]*\n\t// \tctx\.JSON[^\n]*\n\t// } else {\n', '', content)
    content = re.sub(r'\n\t// }', '', content)
    
    # standardize successfull to successful
    content = content.replace('successfull', 'successful')
    
    with open(path, 'w') as f:
        f.write(content)

