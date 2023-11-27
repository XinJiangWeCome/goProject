cookie只是保存在本地
session保存在服务端
token认证：
1、token认证分为三个部分，header, payload, signature,其中header中包含加密方式，payload中含有主要信息，signature则是由header与payload两部分组成
2、token中header由两部分组成，token类型+签名算法
3、payload中主要是一些明文信息，用户名时间戳等信息（注意其中的信息是公开的，不要把重要信息放payload内）
4、signature由三部分组成，header+编码过的payload+服务器上面的密钥+header中的签名算法
JWT安全性的保障
1、不要把隐私存在payload中
2、密钥是核心，一定不要泄露出去。签名的核心就是密钥
3、payload一定要添加过期时间