1.功能
------
	根据数据库表的描述文件自动生成go代码文件和sql文件

2.示例
------
	数据库表的描述文件 example.table:
	------------- file begin --------------
	/****************
	 * 先指明数据库名称
	 */
	database MyDB;

	// 定义一个名为 Test 的表
	struct Test {
		int key; // 域申明以分号结束
		int propery;
		varchar30 name;
	}; // 花括号结尾需要分号，与c语法一致

	// 定义一个名为 People 的表
	struct People {
		int64 id; // 支持 int,int8,int16,int32,int64 等 5 种整数类型
		string name; // string = varchar30
		int8 sex;
		int16 age;
	};
	------------- file end  ---------------

3.说明
------
(1)关键字：database,struct,int,int8,int16,int32,int64,float,double,string, 以及一个后缀可变的类型关键字varchar<n>(如varchar5,varchar30).

(2)表描述文件应首先指明数据名称(如 database MyDB;),然后定义各个表，语法完全与c的结构体定义一致.

(3)支持c++风格的注释

