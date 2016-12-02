#pragma once

#include "LcTypes.h"
using namespace leetcode;

namespace leetcode {
	class Solution {
	public:
		/* Suppose we dont need the original l1,l2 list after the sum. */
		ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
			int ex = 0;
			// new header node for the sum result
			ListNode* lsum = new ListNode(0);
			ListNode* ltmp = lsum;
			lsum->next = l1;
			while (l1 != nullptr && l2 != nullptr) {
				// do sum
				int sum = l1->val + l2->val + ex;
				l1->val = sum % 10;
				ex = sum / 10;
				// move forward
				l1 = l1->next;
				auto unused = l2; // l2 node has been added to l1, so it's unused any more.
				l2 = l2->next;
				delete unused;
				ltmp = ltmp->next;
			}

			if (l1 == nullptr) {
				ltmp->next = l2;
			}

			delete l1, l2; // suppose we dont need the original l1,l2 list

			while (ex >0) {
				if (ltmp->next == nullptr) {
					ltmp->next = new ListNode(ex);
					break;
				}

				int sum = ltmp->next->val + ex;
				ltmp->next->val = sum % 10;
				ex = sum / 10;

				ltmp = ltmp->next;
			}

			return lsum->next;
		}

	};
}
