/**
 *@Author luojunying
 *@Date 2022-01-09 23:04
 */
package main

import "sync"

type singleton struct{}

var ins *singleton

var mt sync.Mutex

//1.升级版1
func NewSingleton() *singleton {
	if ins == nil {
		mt.Lock()
		defer mt.Unlock()
		if ins == nil {
			ins = &singleton{}
		}
	}
	return ins
}

//优点: 没必要一进来就加锁, 一定程度提交高了并发
//缺点: 还是会有网络拥塞的情况, 如果实例化比较久

//1.升级版2
var once sync.Once

func NewSingleton2() *singleton {
	once.Do(func() {
		ins = &singleton{}
	})
	return ins
}

//理论与上面第一个升级版 相差无几

/**
原文链接：https://www.cnblogs.com/stone94/p/10409669.html
原子性：一个事务内的操作，要么同时成功，要么同时失败
一致性：一个事务必须使数据库从一个一致性状态变换到另一个一致性状态

对于一致性，知乎链接中内容如下：

从这段话的理解来看，所谓一致性，即，从实际的业务逻辑上来说，最终结果是对的、是跟程序员的所期望的结果完全符合的

重点
一致性是基础，也是最终目的，其他三个特性（原子性、隔离性和持久性）都是为了保证一致性的
在比较简单的场景（没有高并发）下，可能会发生一些数据库崩溃等情况，这个时候，依赖于对日志的 REDO/UNDO 操作就可以保证一致性
而在比较复杂的场景（有高并发）下，可能会有很多事务并行的执行，这个时候，就很可能导致最终的结果无法保证一致性

即，这个时候，原子性不能保证一致性。因为从单个事务的角度看，不管是事务 1 还是事务 2，它们都保证的原子性（单个事务内的所有操作全部成功了），但最终，它们并没有保证数据库的一致性（因为从逻辑上说，账户 A 应该增加了 200 元，而不是 100 元）
所以，为了保证并发情况下的一致性，又引入了隔离性的概念

隔离性：即事务之间感知不到彼此的存在，就好像只存在本身一个事务一样
而对于怎样实现隔离性，又涉及到了乐观锁和悲观锁的概念

小小引申：
不考虑隔离性的时候，可能导致脏读、幻读和不可重复读的问题（这些问题，其实就是导致无法保证一致性的几种情况）
而隔离级别的概念，就是为了解决上述三个问题
*/
