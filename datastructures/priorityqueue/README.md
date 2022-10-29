# What is a Priority Queue?
 
A priority queue is an Abstract Data Type (ADT) that operates similar to a normal queue
except that **each element has a certain priority**. The priority of the elements in the 
priority queue determine the order in which elements are removed from the PQ. 

*NOTE*: Priority queues only supports comparable data, meaning the data inserted into 
the priority queue must be able to ordered in some way either from least to greatest or
greatest to least. This is so that we are able to assign relative priorities to each element. 

# What is a Heap?

A heap is a tree based Data Structure t hat satisfies the heap invariant (also called heap property):
If A is a parent node of B then A is ordered with respect to B for all nodes A,B in the heap.

I.e.\
*The value of the parent node is always greater than or equal to the value of child node for all nodes* (Max Heap).\
*The value of the parent node is always less than or equal to the value of the child node for all nodes* (Min Heap).

Heaps forms the canonical underlying data structure for PQ. Heaps doesn't contain any cycles. 

There are many types of heaps we could use to implement a priority queue including:
- Binary heap
- Fibonacci Heap
- Binomial Heap 
- Pairing Heap

# Priority Queue With Binary Heap 

A binary heap is a binary tree that supports the heap invariant. In a binary 
tree very node has exactly two children. 

A complete binary tree is a tree which at every level, except possibly the last is 
completely filled and all the nodes are far left as possible. 

Maintain the complete binary tree property is very important because it gives us an insertion point no matter
what the heap looks like or what values are in it. 

Array is a canonical way for representing binary heaps. Using an array is convenient because 
maintaining the complete tree property the insertion position is just the last position in our array. 
Heaps also can be represented by using objects and pointers.

Storing binary heap in array allows easily access  all the children and parent nodes. 

Let i be the parent node index. 
- Left child index: 2i + 1
- Right child index: 2i + 2

# When and where is a priority queue used?

1) Uses in certain implementations of Dijkstra's Shortest Path algorithm.
2) Anytime you need the dynamical fetch 'next best' or 'next worst' element.
3) Used in Huffman coding (which is used for losses data compression).
4) Best First Search (BFS) algorithms such as A use a PQs to continuously grab the next most promising node.
5) Used by Minimum Spanning Tree (MST) algorithms.

**Complexity Analysis**:

| Operation                                      | Time      |
|------------------------------------------------|-----------|
| Binary Heap construction                       | O(n)      |
| Polling                                        | O(log(n)  |
| Peeking                                        | O(1)      | 
| Adding                                         | O(log(n)) |
| Naive Removing                                 | O(n)      |
| Advanced removing with help from a hash table* | O(log(n)) |
| Naive contains                                 | O(n)      |
| Contains check with help of hash table*        | O(1)      |

*Using a hash table to help optimize these operations does take up 
linear space and also adds some overhead to the binary heap implementation. 
