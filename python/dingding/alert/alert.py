import json
import requests

# for python3
def alert(dingUrl :str,msg :str):
    headers = {
        "Content-Type": "application/json"
    }
    data = {
        "msgtype": "text",
        "text": {
            "content": "alert " + msg
        }
    }
    res = requests.post(
        url=dingUrl,
        data=json.dumps(data), headers=headers)
    print(res.status_code)

if __name__ == '__main__':
    alert("url1","test")

