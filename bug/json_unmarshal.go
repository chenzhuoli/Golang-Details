"encoding/json" //json反序列化的天坑，坑死老娘了

type DataType struct { 
    name string 
} 

var data *DataType
err := json.Unmarshal([]byte(dataBytes), data) // 此时无论如何都会报错

//正确的用法是: 
var data DataType //不要用指针 
err = json.Unmarshal([]byte(dataBytes), &data) //反序列化成功,一点点的不同结果千差万别
