1、一般来说，我们通过创建 struct 来实现接口。然而，struct 的用途是用于存储数据，但是目前没有状态可存储，因此创建一个 struct 感觉不太对。
# 接口通过struct的方法实现，struct的属性用于存储数据，没有状态可存储，可以通过其他类型的方法实现接口

2、代码重构，得分计算从 handler 移到函数 GetPlayerScore 中，这就是使用接口重构的正确方法。把重构的函数改成接口。
# 重构方向 把代码为重构为特定函数，把函数改成接口

3、为了让 PlayerServer 能够使用 PlayerStore，它需要一个引用。现在是改变架构的时候了，将 PlayerServer 改成一个 struct。
# 重构方向 函数想使用接口，可以将函数改为结构体，接口成为结构体中的引用