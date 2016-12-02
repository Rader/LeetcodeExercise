#pragma once

#include "LcTypes.h"
#include <stack>
using namespace leetcode;
using namespace std;

namespace leetcode {
	class Solution {
	public:
		/*
		1.The most significant digit comes first and each of their nodes contain a single digit. 		
		2.Cannot modify the input lists.In other words, reversing the lists is not allowed.
    
		Example:

		Input: (7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
		Output: 7 -> 8 -> 0 -> 7
    
    Solution: Push nodes into a stack and then pop them one by one, we get the least significant digits first. 
    then we can add numbers as normal. The result nodes push to a stack, and then we read them again, to build the result list.
    So we get the most significant digits first.
		*/
		ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
			auto s1 = toStack(l1);
			auto s2 = toStack(l2);
			auto ssum = new stack<ListNode*>();

			int ex = 0;
			while (!s1->empty() || !s2->empty() || ex > 0)
			{
				int sum = ex;
				if (!s1->empty()) {
					sum += s1->top()->val;
					s1->pop();
				}
				if (!s2->empty()) {
					sum += s2->top()->val;
					s2->pop();
				}

				ssum->push(new ListNode(sum % 10));// save sum result
				ex = sum / 10;
			}

			delete s1, s2;
			auto lsum = toList(ssum);
			delete ssum;
			return lsum;
		}

		stack<ListNode*>* toStack(ListNode* l1) {
			stack<ListNode*>* s = new stack<ListNode*>();
			while (l1!=nullptr)
			{
				s->push(l1);
				l1 = l1->next;
			}

			return s;
		}

		ListNode* toList(stack<ListNode*>* s) {
			auto lsum = new ListNode(0);
			auto ltmp = lsum;
			while (!s->empty())
			{
				ltmp->next = s->top();
				ltmp = ltmp->next;
				s->pop();
			}

			return lsum->next;
		}

	};
}
