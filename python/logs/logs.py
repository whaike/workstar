# coding:utf-8
import logging.handlers

# logging初始化工作
logging.basicConfig()
# myapp的初始化工作
logs = logging.getLogger('logs')
logs.setLevel(logging.INFO)

# 指定logger输出格式
formatter = logging.Formatter('%(asctime)s %(filename)s [line:%(lineno)d] %(funcName)s %(levelname)s %(message)s %(process)d ')

# 写入文件，如果文件超过N个Bytes(这里设置100M)，仅保留5个文件。
handler = logging.handlers.RotatingFileHandler('test.log', maxBytes=104857600, backupCount=5)
handler.setFormatter(formatter)
logs.addHandler(handler)


def main():
	logs.info("test001")
	logs.error("test002")


if __name__ == '__main__':
	main()