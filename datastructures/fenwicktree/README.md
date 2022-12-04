# What is a Fenwick Tree?

A Fenwick tree (also called Binary Indexed Tree) is a data structure that supports sum range queries as well as 
setting values in a static array and getting the value of the prefix sum up some index efficiently. Fenwick tree
is a one based. 

Unlike a regular array, in a Fenwick tree a specific cell is responsible for other cells as well. 

The position of the least significant bit (LSB) determines the range of responsibility that cell has to the cells
below itself. 

All odd numbers have a first least significant bit set in the ones position, so they are only responsible for 
themselves. 

In a Fenwick tree we may compute the prefix sum up to a certain index, which ultimately lets us perform range 
sum queries.  

**Idea**: Suppose you want to find the prefix sum of [1, i], then you **start at i and cascade downwards** until you reach
zero adding the value at each of the indices you encounter. 

Compute the interval sum between [11, 15]. First we compute the prefix sum of [1, 15], then we will compute the prefix 
sum of [1, 11) and get the difference. 

Notice that in the worst case the cell we're querying has a binary representation of all ones (numbers of the form
2^n - 1).

Hence, it's easy to see that in the worst case a range query might make us have to do two queries that cost log2(n)
operations. 

Range query algorithm

To do a range query from [i, j] both inclusive a Fenwick tree of size N:

```
function pefixSum(i):
    sum := 0
    while i != 0:
        sum = sum + tree[i]
        i = i - LSB(i)
    return sum
    
function rangeQuery(i, j):
    return prefixSum(j) - prefixSum(i - 1) 
```

Where **LSB** returns the value of the least significant bit. 


## Point Updates

Point updates are opposite to cascade down approach presented in range queries. 
Instead, we'd like to add the LSB to propagate the value up the cells responsible for us. 

To update cell at index i in the Fenwick tree of size N:

```
function add(i, x):
    while i < N:
        tree[i] = tree[i] + x
        i = i + LSB(i)
```

Where **LSB** returns the value of the least significant bit. For example:

LSB(12) = 4 because 12 = 1100 and the least significant bit of 1100 is 1**1**00 that equals 4 (2^2)
in base 10.

## Naive construction

Let A be an array of values. For each element in A at index i do a point update on the Fenwick 
tree with a value of A[i]. There are n elements and each point update takes O(log(n)) for a total of
O(nlogn(n)), can we do better?

## Linear construction 

Idea: Add the value in the current cell to the **immediate cell** that is responsible for us. This 
resembles what we did for point updated but only one cell at a time. 

This will make the 'cascading' effect in range queries possible by propagating the value in each cell
throughout the tree. 

Let i be the current index. The immediate cell above us is at position j given by:

j := i + LSB(i)

Where LSB is the **Least Significant Bit** of i. Ignore updating j if index is out of bounds. 

```
# Make sure values is 1-based!
function consturct(values):
    N := length(values)
    
    # Clone the values array since we're 
    # doing in place operations 
    tree = deepCopy(values)
    
    for i = 1,2,3, ... N:
        j := i + LSB(i)
        if j <= N:
            tree[j] = tree[j] + tree[i] 
    return tree 
```

**Complexity Analysis**:

| Operation      | Avg       |
|----------------|-----------|
| Construction   | O(n)      |
| Point Update   | O(log(n)) |
| Range Sum      | O(log(n)) |
| Range Update   | O(log(n)) |
| Adding Index   | N/A       |
| Removing Index | N/A       |
