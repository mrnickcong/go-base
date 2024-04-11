# 记录java转GO需要注意的地方

- [记录java转GO需要注意的地方](#记录java转go需要注意的地方)
  - [一、`&` 和 `*`](#一-和-)
    - [`&` 取地址符](#-取地址符)
    - [`*` 指针](#-指针)
  - [二、函数参数的传递](#二函数参数的传递)
  - [三、`*` 另一个作用：解引用](#三-另一个作用解引用)
  - [四、GO中的函数](#四go中的函数)
  - [五、Go的并发](#五go的并发)
  - [问题](#问题)
    - [Q1：为什么协程块？](#q1为什么协程块)
    - [Q2：协程是不是等同于线程池的任务？](#q2协程是不是等同于线程池的任务)
  - [GO里的缺点](#go里的缺点)

## 一、`&` 和 `*`

### `&` 取地址符

GO语言中 `&` 只有一个含义：获取地址的值

```text
var i int = 8
fmt.Println(&i)
//打印结果
0xc00000e088
```

### `*` 指针

指针就是一个地址的值，指针的类型可以是基本数据类型，也可以是自定义类型的任何类型

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

在java中如果需要一个整数的指针：`Integer i = new Integer()` i就是一个整数类型的指针

## 二、函数参数的传递

java中的函数传递

```text
    public static void modifyName(Person person){
        person.setName("李四");
    }

    public static void main(String[] args) {
        Person person = new Person();
        person.setName("张三");
        modifyName(person);
        System.out.println(person.getName());
        //输出值为：李四
    }
```

GO语言中：只有值传递。包括：结构体的值和指针的值

```text
//传递的是整个Person的结构体的值。复制了一份，传递过去。
func m1(person Person) {
    person.name = "李四"
}

//传递的是值，操作的是同一个地址上的数据
func m2(person *Person) {
    person.name = "李四"
}
```

## 三、`*` 另一个作用：解引用

```text
func mf(m *int) {
//解引用 给变量m赋值
*m = 8
}
```

函数的指针：是只读的，不可改变

## 四、GO中的函数

[见T03-function]

## 五、Go的并发

关键字：go

go 启动协程的原理

GMP模型 goroutine processor machine
go就

## 问题

### Q1：为什么协程块？

### Q2：协程是不是等同于线程池的任务？

类似于java中线程池的task，不具备协调能力

golang可以用channel来做协程协调，在用户空间模拟了CPU的切换

协程在执行完任务后回收

golang中没有异步。因为到处都是异步，虽然是同步的写法

## GO里的缺点

- 没有切面编程，不支持动态代理
- 生态不好
- 没有泛型
