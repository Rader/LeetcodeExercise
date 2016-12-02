#pragma once

namespace leetcode {
	//Definition for singly-linked list.
	struct ListNode {
		int val;
		ListNode *next;
		ListNode(int x) : val(x), next(nullptr) {}
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
