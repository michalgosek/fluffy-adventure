# What is Union Find?

Union find is a data structure that keeps track of elements which are 
split into one or more disjoints sets. Its has two primary operations 
find and union. 

Find tells what group the element belongs to.\
Union merges two groups together. 

# When and where is a Union Find used?

1) Kruskal's minimum spanning tree algorithm
2) Grid percolation
3) Network connectivity
4) Least common ancestor in trees 
5) Image processing 

# Union Find application: Kruskal's Minimum Spanning Tree

Given a graph G = (V, E), we want to find a *Minimum Spanning Tree*
in the graph (it may not be unique). A minimum spanning tree is a 
subset of the edges which connect all vertices in the graph with 
the minimal total edge cost. 

1) Sort edges by ascending edges weight 
2) Walk through the sorted edges and look at two nodes the edge belongs to
,if the nodes are already unified we don't include this edge, otherwise we 
include it and unify the nodes. 
3) The algorithm terminates when every edge has been processed or all the 
vertices have been unified. 

# Creating Union Find
To begin using Union Find, first construct a *bijection* (a mapping)
between your objects and the integers in the range [0, n).

**Note**: This step is not necessary in general, but it will allow us to construct
an array-based union find. 

# Find Operation
To find which component a particular element belongs to find the
root of that component by following the parent nodes until a self 
loop is reached (a node who's parent is itself)

# Union Operation
To unify two elements find which are the root nodes of each component
and if the root nodes are different make one of the root nodes be the
parent of the other. 

# Remarks
In this data structure, we do not "un-union" elements. In general, this
would be very inefficient to do so since we would have to update all the 
children of a node.

The number of components is equal to the number of roots remaining. 
Also, remark that the number of root nodes never increases. 

**Complexity Analysis**: 

| Operation          | Time |
|--------------------|------|
| Construction       | O(n) |
| Union              | α(n) |
| Find               | α(n) |
| Get component size | α(n) |
| Check if connected | α(n) |
| Count components   | O(1) |

α(n) - Amortized constant time. 
