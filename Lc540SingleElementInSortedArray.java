/**
Problem URL: https://leetcode.com/problems/single-element-in-a-sorted-array/description/
Given a sorted array consisting of only integers where every element appears twice except for one element which appears once. Find this single element that appears only once.

Example 1:
Input: [1,1,2,3,3,4,4,8,8]
Output: 2
Example 2:
Input: [3,3,7,7,10,11,11]
Output: 10
Note: Your solution should run in O(log n) time and O(1) space. 
 */

/**
 * @author Rader
 *
 */
public class Lc540SingleElementInSortedArray {
	public int findSingle(int[] vals) {
		int begin = 0;
		int end = vals.length - 1;
		int idx = binarySearch(begin, end, vals);

		return vals[idx];
	}

	private int binarySearch(int begin, int end, int[] vals) {
		if (begin == end) {
			return begin;
		}

		int mid = (begin + end) / 2;
		// System.out.printf("%d,%d,%d, %s\n", begin, mid, end,
		// Arrays.toString(Arrays.copyOfRange(vals, begin, end + 1)));

		if (vals[mid] == vals[mid + 1]) {
			if ((mid - begin) % 2 == 1) { // left part has odd number of elements
				return binarySearch(begin, mid - 1, vals);
			} else {
				// search in right part
				return binarySearch(mid + 2, end, vals);
			}
		} else if (vals[mid] == vals[mid - 1]) {
			if ((end - mid) % 2 == 1) { // right part has odd number of elements
				return binarySearch(mid + 1, end, vals);
			} else {
				// search in left part
				return binarySearch(begin, mid - 2, vals);
			}
		} else {
			// we found it
			return mid;
		}
	}
}
