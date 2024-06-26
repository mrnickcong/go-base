# Go基础部分

- [Go基础部分](#go基础部分)
  - [1.数据类型](#1数据类型)
    - [1.1 数字类型](#11-数字类型)
      - [1.1.1 整形](#111-整形)
      - [1.1.2 浮点](#112-浮点)
      - [1.1.3 其他数字类型](#113-其他数字类型)
    - [1.2 派生类型](#12-派生类型)
  - [运算符](#运算符)
  - [流程控制](#流程控制)
  - [关键字](#关键字)
  - [数组](#数组)
  - [切片](#切片)
  - [映射](#映射)

## 1.数据类型

在 Go 编程语言中，数据类型用于声明函数和变量。

数据类型的出现是为了把数据分成所需内存大小不同的数据，编程的时候需要用大数据的时候才需要申请大内存，就可以充分利用内存。

Go主要有：布尔类型、数字类型、字符串类型、派生类型

### 1.1 数字类型

整型 int 和浮点型 float32、float64，Go 语言支持整型和浮点型数字，并且支持复数，其中位的运算采用补码。

#### 1.1.1 整形

|序号|类型 |描述|长度|
|:---:|:---:|:---:|:---:|
|1|uint8|无符号 8 位整型|0-255|
|2|uint16|无符号 16 位整型||
|3|uint32|无符号 32 位整型||
|4|uint64|无符号 64 位整型||
|5|int8|带符号 8 位整型|-128-127|
|6|int16|带符号 16 位整型||
|7|int32|带符号 32 位整型||
|8|int64|带符号 64 位整型||

#### 1.1.2 浮点

|序号|类型 |描述|长度|
|:---:|:---:|:---:|:---:|
|1|float32|32位浮点型数||
|2|float64|64位浮点型数||
|3|complex64|32 位实数和虚数||
|4|complex128|64 位实数和虚数||

#### 1.1.3 其他数字类型

|序号|类型 |描述|长度|
|:---:|:---:|:---:|:---:|
|1|byte|类似 uint8||
|2|rune|类似 int32||
|3|uint|32 或 64 位||
|4|int|与 uint 一样大小||
|5|uintptr|无符号整型，用于存放一个指针||

### 1.2 派生类型

- 指针类型（Pointer）
- 数组类型
- 结构化类型(struct)
- Channel 类型
- 函数类型
- 切片类型
- 接口类型（interface）
- Map 类型

## 运算符

## 流程控制

## 关键字

## 数组

## 切片

## 映射
