import requests
import ujson
import base64
import cv2
import pprint
import io
import numpy as np

register_url = "http://127.0.0.1:12580/talkuhulk/aidb/register"
unregister_url = "http://127.0.0.1:12580/talkuhulk/aidb/unregister"
query_url = "http://127.0.0.1:12580/talkuhulk/aidb/query"

register_body = {"flow_uuid": "i'm uuid", "model": ["scrfd_500m_kps", "pfpld"], "backend": ["mnn", "mnn"], "zoo": "./config"}
unregister_body = {"flow_uuid": "i'm uuid"}

test_image = "./test.jpg"
with open(test_image, "rb") as f:
    images_base64 = base64.b64encode(f.read()).decode('utf8')

query_body = {"flow_uuid": "i'm uuid", "image_base64": images_base64}

request_body = ujson.dumps(query_body)

response = requests.post(url=query_url, data=request_body, headers={"content-type": "application/json"})
if not response.ok:
    print(response.text)
    exit(-1)
result = response.json()

if result["code"] == 0:
    bgr = cv2.imread(test_image)
    if result["face"]:
        for face in result["face"]:
            bgr = cv2.rectangle(bgr,
                                (int(face["bbox"][0]), int(face["bbox"][1])),
                                (int(face["bbox"][2]), int(face["bbox"][3])),
                                (0, 0, 255), 2)
            for x, y in face["landmark"]:
                bgr = cv2.circle(bgr, (int(x), int(y)), 1, (255, 0, 0), -1)
    if result["tddfa"]:
        images_raw = base64.b64decode(result["tddfa"])
        io_buf = io.BytesIO(images_raw)
        decode_tddfa = cv2.imdecode(np.frombuffer(io_buf.getbuffer(), np.uint8), -1)
        cv2.imshow("tddfa", decode_tddfa)
    pprint.pprint(result)
    cv2.imshow("result", bgr)

    cv2.waitKey()
else:
    print(result)
    print("failed~")


