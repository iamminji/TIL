## 파이썬에서 시간 복잡도


기존에 스택이나 큐를 써야되는 일이 있으면 왠만해서는 <code>collections.deque</code>을 사용했었다.

그러다가, 우연히 스택 처럼 리스트를 쓰려면 list.pop(0)을 하면 된다는 이야기를 듣고 꾸준히 그렇게 썼었는데 알고보니까 list.pop(index) 는  O(N) 이었다.

list.pop() 은 당연히 O(1) 이고.


https://www.ics.uci.edu/~pattis/ICS-33/lectures/complexitypython.txt
https://groups.google.com/forum/#!topic/comp.lang.python/kiHYf5N0iz8
https://stackoverflow.com/questions/195625/what-is-the-time-complexity-of-popping-elements-from-list-in-python
