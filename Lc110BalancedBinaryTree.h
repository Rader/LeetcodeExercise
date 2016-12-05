#pragma once

#include <algorithm>
#include <cassert>
#include "LcTypes.h"

using namespace std;

namespace leetcode {
	class Solution {
	public:
		bool isBalanced(TreeNode* root) {
			if (root == nullptr)
				return true;

			auto leftHeight = heightOfTree(root->left);
			auto rightHeight = heightOfTree(root->right);
			
			if (abs(leftHeight - rightHeight) > 1)
				return false;

			return isBalanced(root->left) && isBalanced(root->right);
		}

		int heightOfTree(TreeNode* root) {
			if (root == nullptr)
				return 0;

			int height = 1;
			return height + max(heightOfTree(root->left), heightOfTree(root->right));
		}

		void isBalancedTest() {
			TreeNode* root = new TreeNode(1);
			root->left = new TreeNode(2);
			root->right = new TreeNode(3);
			root->left->left = new TreeNode(4);
			root->left->left->left = new TreeNode(5); // unbanlanced

			assert(isBalanced(root));
		}
	};
}
