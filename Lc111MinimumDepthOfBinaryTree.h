#pragma once

#include "LcTypes.h"
#include <algorithm>
#include <cassert>

namespace leetcode {
	class Solution {
	public:
    /* Need to take care of a tree with has only one subtree, left or right subtree is empty.*/
		int minDepth(TreeNode* root) {
			if (root == nullptr)
				return 0;

			// if one of the subtrees is empty, return the depth of another subtree
			if (root->left == nullptr) return 1 + minDepth(root->right);
			if (root->right == nullptr) return 1 + minDepth(root->left);

			return 1 + std::min(minDepth(root->left), minDepth(root->right));
		}

		void minDepthTest()
		{
			TreeNode* root = new TreeNode(1);
			root->left = new TreeNode(2);
			root->right = new TreeNode(3);
			root->left->left = new TreeNode(4);
			root->left->left->left = new TreeNode(5);

			assert(2 == minDepth(root));

			root->left = nullptr; // Tree only has right subtree: 1->3
			assert(2 == minDepth(root));
		}
	};
}
