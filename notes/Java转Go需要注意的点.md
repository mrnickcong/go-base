# 记录java和GO的区别

## 一、`&` 和 `*`

### `&` 取地址符

GO语言中 `&` 只有一个含义：获取地址的值

```text
var i int = 8
fmt.Println(&i)
//打印结果
0xc00000e088
```

## `*` 指针

指针就是一个地址的值，指针的类型可以是基本数据类型，也可以是自定义类型

```text
var i int = 8
fmt.Println(&i)

//pi是一个指向指针类型的指针
var pi *int 
//指针赋值
pi  = &i
fmt.Println(pi)
//打印结果
0xc00000e088
0xc00000e088
```

https://www.mashibing.com/study?courseNo=374&sectionNo=147&courseVersionId=1283