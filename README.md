# Print Order Address Label with QL-800

!! DEPRECATED !! merged to [thermal-station](https://github.com/suapapa/thermal-station)

```bash
curl -X POST \
-d '{"ID":"1234567890","from":{"line1":"경기 성남시 분당구 판교역로 235 (에이치 스퀘어 엔동)","line2":"7층","name":"카카오 엔터프라이즈","phone_number":"010-1234-5678"},"to":{"line1":"경기도 성남시 분당구 판교역로 166","name":"판교 아지트","phone_number":"010-1234-5678"}}' \
http://rpi-airplay.local:8080/v1/order
```
