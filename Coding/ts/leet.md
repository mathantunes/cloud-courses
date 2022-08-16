# Coding 4 fun on Typescript

## Contains Duplicate

Creates an array of uniques and compares the lenght
```typescript
function containsDuplicate(nums: number[]): boolean {
    const uniques = [...new Set(nums)];
	// Compares length of a set of uniques and the input set
    return uniques.length !== nums.length;
};
```

## Contains Nearby Duplicate

Given an integer array nums and an integer k, 
return true if there are two distinct indices i and j in the array such that nums[i] == nums[j] 
and abs(i - j) <= k.

Time: O(N) -> from the some call

```typescript
function containsNearbyDuplicate(nums: number[], k: number): boolean {
    // First make sure to only handle arrays with duplicates
    const uniques = [...new Set(nums)];
    if(nums.length === uniques.length) {
        return false;
    }
    // Hash stores [num -> index on nums]
    const hash = {};
    return nums.some((num, idx) => {
        // num in hash -> current num is duplicate
        // Math operation according to the problem restriction
        if(num in hash && Math.abs(idx - hash[num]) <= k) {
            return true;
        }
        hash[num] = idx;
    });
};
```

## Contains Nearby Almost Duplicates

Given an integer array nums and two integers k and t,
return true if there are two distinct indices i and j in the array such that
abs(nums[i] - nums[j]) <= t and abs(i - j) <= k.

In this exercise, we don't have the duplicates to compare. therefore using the slices defined on K helps limiting the loops

Time: O(N) -> from the some call

```typescript
function containsNearbyAlmostDuplicate(nums: number[], k: number, t: number): boolean {
    return nums.some((num, idx) => {
        // Slice the array in steps of K -> abs(i - j) <= k
        const sliced = nums.slice(idx+1, idx+k+1);
        for (let i = 0; i < sliced.length; i++) {
            // Current iteration has to be tested against the neighbors
            if (Math.abs(num - sliced[i]) <= t) {
                return true;
            }
        }
    });
};
```

## Max subarray sum (DYNAMIC)

```typescript
// [-2,1,-3,4,-1,2,1,-5]
// Start with [-2, 1] => Max: 1, TotalMax: 1
// Start with [1, -3] => Max: -2 (1 + (-3)), TotalMax: 1 (preserved)
// Start with [-3, 4] => Max: 4, TotalMax: 4
// Start with [4, -1] => Max: 3, TotalMax: 4 (preserved)
// Start with [-1, 2] => Max: 5, TotalMax: 5
// Start with [2, 1] => Max: 6, TotalMax: 6
// Start with [1, -5] => Max: 1, TotalMax: 6 (preserved)
function maxSubArray(nums: number[]): number {
    let max = nums[0];
    let totalMax = max;
    // Accumulated + next is compared with next alone
    // Which is then compared with the totalMax
    for(let i = 1; i < nums.length; i++) {
        max = Math.max(max + nums[i], nums[i]);
        totalMax = Math.max(max, totalMax);
    }
    return totalMax;
};
```

## Max subarray product (DYNAMIC)

Handles the negative values differently.
Have to flip currentMax and currentMin and store both to make sure we can flip back if there's another negative to multiply

```typescript
function maxProduct(nums: number[]): number {
    let currentMax = nums[0];
    let totalMax = nums[0];
    let currentMin = nums[0];
    let temp = 1;
    for(let i = 1; i < nums.length; i++) {
        if (nums[i] < 0)
        {
            temp = currentMax;
            currentMax = currentMin;
            currentMin = temp; 
        }
        currentMax = Math.max(currentMax * nums[i], nums[i]);
        currentMin = Math.min(currentMin * nums[i], nums[i]);
        totalMax = Math.max(totalMax, currentMax);
        
        // console.log("a", { num: nums[i], currentMax, totalMax})
    }
    return totalMax;
};
```

## Two Sum

Register indexes already visited to not reuse.
Try to find the element index for the subtraction between current and target

```typescript
function twoSum(nums: number[], target: number): number[] {
    const visited = {};
    const len = nums.length;
    for(let i = 0; i < len; i++) {
        const num = nums[i];
        const rest = target - num;
        const restIdx = nums.indexOf(rest);
        visited[num] = i; // Register visit to this index
        if (visited[rest] !== restIdx && // Index can't be reused
			restIdx > -1) {
            return [i, restIdx];
        }
    }
    return [];
};
```

## Three Sum

```typescript
function threeSum(nums: number[]): number[][] {
    let n = nums.length
    let res = [];
    nums.sort((a,b) => a-b); // Sort array before using

    for (let i = 0; i < n; i++) {
        const a = nums[i]; // First term only looping once
        if (i > 0 && a === nums[i-1]) continue // If repeated, ignore
        let l = i + 1;
        let r = n - 1;
        while(l < r) { // Since it's sorted, try to find the sum which results in zero
            const sum = a + nums[l] + nums[r];
            if (sum > 0) {
                r--; // Too high, lower the end pointer
            } else if (sum < 0) {
                l++; // Too low, higher the start pointer
            } else {
                res.push([a, nums[l], nums[r]]);
				// Continue with start pointer until next value
                l++;
                while(nums[l] === nums[l-1] && l < r) l++;
            }
        }
    }
    return res;
};
```

## Count K Difference

Given an integer array nums and an integer k, return the number of pairs (i, j) where i < j such that |nums[i] - nums[j]| == k.

```typescript
function countKDifference(nums: number[], k: number): number {
    const visited = {};
    let counter = 0;
    for (let i = 0; i < nums.length; i++) {
        const num = nums[i];
		// If the rest has already been visited, add to the counter
        if(visited[num - k]) {
            counter += visited[num - k]
        }
        if(k !== 0 && visited[num + k]) {
            counter += visited[num + k]
        }
		// Create a counter for number on the sequence
        visited[num] = visited[num] ? visited[num] + 1 : 1;
    }
    return counter;
};
```

## Merge sorted array

```typescript
function merge(nums1: number[], m: number, nums2: number[], n: number): void {
    /*
        [1,2,3]
        [2,3,4]
        -> last = 5, m = 3, n = 3 -> nums1 = [1,2,3,0,0,0]
        -> last = 4, m = 3, n = 2 -> nums1 = [1,2,3,0,0,4]
        -> last = 3, m = 3, n = 1 -> nums1 = [1,2,3,0,3,4]
        -> last = 2, m = 2, n = 1 -> nums1 = [1,2,3,3,3,4]
        -> last = 1, m = 2, n = 0 -> nums1 = [1,2,2,3,3,4]
        (stop because n = 0, num1 elements are already sorted)
    */
    let last = m + n - 1;
    // Merge in reverse -> nums1 ends with N zeroes
    while(m > 0 && n > 0) {
        if (nums1[m-1] > nums2[n-1]) {
            nums1[last] = nums1[m-1]
            m--;
        } else {
            nums1[last] = nums2[n-1]
            n--;
        }
        last--;
    }
    // fill leftovers of nums2
    while (n > 0) {
        nums1[last] = nums2[n-1]
        n--;
        last--;
    }
};
```

## Intersect

With sorting, we can use pointers and increment them according to the underlying value

```typescript
function intersect(nums1: number[], nums2: number[]): number[] {
    let n1 = 0;
    let n2 = 0;
    let res = [];
    nums1.sort((a,b) => a-b);
    nums2.sort((a,b) => a-b);
    while(n1 < nums1.length && n2 < nums2.length) {
        if(nums1[n1] === nums2[n2]) {
            res.push(nums1[n1]);
            n1++;
            n2++;
        } else if (nums1[n1] > nums2[n2]) {
            n2++;
        } else {
            n1++;
        }
    }
    return res;
};
```

## Add Two Numbers

```typescript
/**
 * Definition for singly-linked list.
 * class ListNode {
 *     val: number
 *     next: ListNode | null
 *     constructor(val?: number, next?: ListNode | null) {
 *         this.val = (val===undefined ? 0 : val)
 *         this.next = (next===undefined ? null : next)
 *     }
 * }
 */

function addTwoNumbers(l1: ListNode | null, l2: ListNode | null): ListNode | null {
    const l1Num = extractAllDigits(l1).split('').reverse();
    const l2Num = extractAllDigits(l2).split('').reverse();
	// Have to do the sum of the reversed values beforehand to get the sum right
    const sum = BigInt(l1Num.join('')) + BigInt(l2Num.join(''));
    const returnDigits = sum.toString().split('').reverse();
	// Create the Linked List with the values
    let ret = new ListNode();
    let next = ret;
    for(let i = 0; i < returnDigits.length; i++) {
        const n = returnDigits[i];
        next.val = Number(n);
        if (i < returnDigits.length -1) {
            next.next = new ListNode();
            next = next.next;
        }
    }
    return ret;
};

function extractAllDigits(l: ListNode | null): string {
    let acc = '';
    while (l !== null) {
        acc = acc + l.val.toString();
        l = l.next;
    }
    return acc;
}
```

## Best time to buy and sell stock

```typescript
function maxProfit(prices: number[]): number {
    let buy = prices[0]; // Starting buy
    let sell = 0;
    let profit = 0;
    const len = prices.length;
    for(let i = 0; i < len; i++) {
       const p = prices[i];
       buy = Math.min(buy, p);
       if (buy === p) {
           // Reset sell if buy updates
           sell = 0;
       } else {
           sell = Math.max(sell, p);
       }
       profit = Math.max(profit, sell - buy, 0);
    }
    return profit;
};
```

## Best time to buy and sell stock II

```typescript
function maxProfit(prices: number[]): number {
    let profit = 0;
    for(let i = 1; i < prices.length; i++) {
        // D(0) - D(-1) -> Defines daily profit
        // If there's any profit, sum it up
        const currentProfit = prices[i] - prices[i - 1];
        profit = profit + (currentProfit > 0 ? currentProfit : 0);
    }
    return profit;
};
```

## Remove duplicates

```typescript
function removeDuplicates(s: string): string {
    const letters = s.split('');
    const stack = [];
    let stackIdx = 0;
    for(let i = 0; i < letters.length; i++) {
       if(stackIdx > 0 && stack[stackIdx - 1] === letters[i]) {
           stackIdx--;
       } else {
           stack[stackIdx] = letters[i];
           stackIdx++;
       }
    }
    return stack.slice(0, stackIdx).join('');
};
```

## Final Prices

```typescript
function finalPrices(prices: number[]): number[] {
    const stack = []; // Stores the indexes that could receive a discount
    for(let i = 0; i < prices.length; i++) {
        const current = prices[i];
         while (stack.length > 0 && // Discount to all items on the stack
		 	prices[stack[stack.length -1]] >= current
		) {
             prices[stack.pop()] -= current; // Pop the index and modify the value on prices
         }
        stack.push(i);
    }
    return prices;
};
```

## Majority Element

```typescript
function majorityElement(nums: number[]): number {
    let candidate = 0;
    let count = 0;
    for (let i = 0; i < nums.length; ++i) {
        if (count == 0) candidate = nums[i];
        count += candidate == nums[i] ? 1 : -1;
    }
    return candidate;
};

/*
Boyer–Moore majority vote algorithm
[2,2,1,1,1,2,2] 
0 -> candidate = 2 count = 1
1 -> candidate = 2 count = 2
2 -> candidate = 2 count = 1
3 -> candidate = 2 count = 0
4 -> candidate = 1 count = 1
5 -> candidate = 1 count = 0
6 -> candidate = 2 count = 1
*/
```

## Merge Intervals

```typescript
// an interval has only 2 values (min, max)
function merge(intervals: number[][]): number[][] {
    intervals.sort((a,b) => a[0] - b[0]); // Sort it before to start with the lowest
    let min = intervals[0][0];
    let max = intervals[0][1];
    let res = [];
    for (let i = 0; i < intervals.length ; i++) {
        const interval = intervals[i];
        if (max < interval?.[0]) { // new min is greater the old max -> no overlapping
            res.push([min, max]);
            min = interval[0]; // new min and max for new interval
            max = interval[1];
        } else {
            max = Math.max(max, interval[1]); // update the max because the min overlaps
        }
    }
    res.push([min, max]);
    return res;
};
```

## Rotate Matrix 90 deg

```typescript
function rotate(matrix: number[][]): void {
    let n = matrix.length;
    for (var i = 0; i < n/2; i++) { // Only iterates halfway through
        for (var j = i; j < n-i-1; j++) {
            // console.log(matrix);
            var tmp = matrix[i][j];
            matrix[i][j] = matrix[n-j-1][i];
            // console.log('replace', { a: [i,j],b: [n-j-1, i]})
            matrix[n-j-1][i] = matrix[n-i-1][n-j-1];
            // console.log('replace', { a: [n-j-1,i],b: [n-i-1, n-j-1]})
            matrix[n-i-1][n-j-1] = matrix[j][n-i-1];
            // console.log('replace', { a: [n-i-1,n-j-1],b: [j, n-i-1]})
            matrix[j][n-i-1] = tmp;
            // console.log('replace', { a: [j,n-i-1],b: [i, j]})
        }
    }
};
```

### Search sorted matrix

Write an efficient algorithm that searches for a value target in an m x n integer matrix matrix.
This matrix has the following properties:

Integers in each row are sorted in ascending from left to right.
Integers in each column are sorted in ascending from top to bottom.

```typescript
function searchMatrix(matrix: number[][], target: number): boolean {
    for (let i = 0; i < matrix.length; i++) {
        const row = matrix[i];
        if (target <= row[row.length - 1] && target >= row[0]) { // Only go through row if target is in range
            for (let j = 0; j < row.length; j++) {
                if (row[j] === target) {
                    return true;
                }
            }
        }
    }
    return false;
};
```