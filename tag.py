import subprocess

tag = ''
# tag前缀
tag_prefix = "v"
# tag分隔符
tag_sep = "."
# 初始化版本
tag_first_version = "v1.0.0"
# 执行 git tag
process = subprocess.run(['git','tag'],capture_output=True,text=True)
if process.returncode == 0:
    tags = process.stdout.splitlines()
    # 已存在旧标签
    if len(tags) > 0:
        # 获取最后一个tag
        last_tag = tags[-1]
        # 去除前缀`v` 和`.`符号,获取版本号
        last_tag_num = last_tag.replace(tag_prefix,"").replace(tag_sep,"")
        # 生成下一个tag版本号
        new_tag_num = int(last_tag_num)+1
        num_list = list(str(new_tag_num))
        new_tag = tag_sep.join(num_list)
        tag = "%s%s"%(tag_prefix,new_tag)
    # 没有旧标签,初始化第一个标签
    else:
        tag = tag_first_version
    print("%s --> %s"%(last_tag,tag))

    # 生成新的tag
    task = subprocess.run(['git','tag',tag],capture_output=True,text=True)
    if task.returncode == 0:
        print("Crate Tag %s Success"%tag)
        # 推送tag到git仓库
        if subprocess.run(['git','push','origin',tag]).returncode == 0:
            print("Push Tag %s Success"%tag)
else:
    # git tag命令执行失败
    print("git tag命令执行失败")