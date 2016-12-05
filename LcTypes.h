#pragma once

namespace leetcode {
	//Definition for singly-linked list.
	struct ListNode {
		int val;
		ListNode *next;
		ListNode(int x) : val(x), next(nullptr) {}
	};

	/* Definition for a binary tree node.*/
	 struct TreeNode {
	     int val;
	     TreeNode *left;
	     TreeNode *right;
	     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
	 };

	static ListNode* BuildList(int* numbers, int count) {
		ListNode* head = new ListNode(numbers[0]);
		auto tail = head;
		for (int i = 1; i < count; i++)
		{
			tail->next = new ListNode(numbers[i]);
			tail = tail->next;
		}
	
		return head;
	}
}
