# Graphs

## Tree

Tree like connection between nodes using edges

## Adjacency Matrix

Matrix of relationship that defines a value for an edge in a matrix view

## Adjacency List

Object where each item has the reference of it's connections

```typescript
const adjacencyList = new Map();
const addNode = (node) => {
	adjacencyList.set(node, []);
};
const addEdge = (start, end) => {
	adjacencyList.get(start).push(end);
	adjacencyList.get(end).push(start);
};
```

## BFS

```typescript
const bfs = (start, target) => {

	const queue = [start];
	const visited = new Set();

	while(queue.length > 0) {
		const node = queue.shift();
		const edges = adjacencyList.get(node);
		for (const edge of edges) {
			if(edge === target) {
				return target;
			}
			if(!visited.has(edge)) {
				queue.push(edge);
				visited.add(edge);
			}
		}
	}
};
```

## DFS

```typescript
const dfs = (start, target, visited = new Set()) => {
	
	visited.add(start);

	const edges = adjacencyList.get(start);
	for(const edge of edges) {
		if(edge === target) {
			return target;
		}
		if(!visited.has(edge)) {
			dfs(edge, target, visited);
		}
	}
};
```


## Quick sort

```typescript

// start with arr, 0, arr.length - 1

function quickSort(arr, start, end) {
    if (start < end) {
        // The partitioning index is represented by pi.
        let pi = partition(arr, start, end);
        // Separately sort elements before and after partition
        quickSort(arr, start, pi - 1);
        quickSort(arr, pi + 1, end);
    }
}

function partition(arr, start, end) {
    // pivot - last element
    let pivot = arr[end];
    /* Index of a smaller element that specifies the pivot's correct position so far. */
    let pivotIndex = start;
    for (let j = start; j <= end - 1; j++) {
        // If current element is smaller than the pivot
        if (arr[j] < pivot) {           
			swap(arr, pivotIndex, j);
            pivotIndex++;
        }
    }
    swap(arr, pivotIndex, end);
    return (pivotIndex);
}

function swap(arr, i, j) 
{   
	let temp = arr[i];
    arr[i] = arr[j];
    arr[j] = temp;
}
```




