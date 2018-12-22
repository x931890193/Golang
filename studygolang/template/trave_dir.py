import os
import re


def gci(filepath):
    # 遍历filepath下所有文件，包括子目录
    files = os.listdir(filepath)
    for fi in files:
        fi_d = os.path.join(filepath, fi)
        if os.path.isdir(fi_d):
            gci(fi_d)
        else:
            if fi_d.endswith('.html'):
                try:
                    with open(fi_d, 'r') as f:
                        data = f.read()
            #    print(data)
                    new_data = re.sub(r'Go语言中文网 - Golang中文社区', 'BSD 资源分享-技术交流网', data)
                    with open(fi_d, 'w') as g:
                        g.write(new_data)
                    print(new_data)
                except:
                    pass

# 递归遍历/root目录下所有文件

gci('./')


def get_file_name(filepath):
    files = os.listdir(filepath)
    for file in files:
        fi_d = os.path.join(filepath, file)
        if os.path.isdir(fi_d):
            get_file_name(fi_d)
        else:
            print(fi_d)
