### goalng基础

1. interface的理解  
  * `interface`和`struct`在golang中都是一种类型，实现了一个`interface`中所有方法的类型就是该`interface`类型。  
  
  * 所有类型都是含有0个方法的`interface`类型，表示为`interface{}`
  
  *  `interface`类型做为函数参数可以接收任意实现了该`interface`的参数
    
  *  example：
    
    ```go
    // /usr/local/Cellar/go/1.13.1/libexec/src/fmt/print.go
    type Stringer interface {
    	String() string
    }
    ```
    
    ```go
    type Person struct {
    	Name string
    	Age int
    }
    
    func (p Person) String () string{
    	return fmt.Sprintf("姓名：%s, 年龄：%d", p.Name, p.Age)
    }
    
    func main(){
    	p := &Person{
    		Name: "test",
    		Age: 21,
    	}
    
    	fmt.Println(p)
    }
    
    // 姓名：test, 年龄：21
    ```
    
    