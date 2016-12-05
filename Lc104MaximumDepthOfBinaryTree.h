#pragma once

#include "LcTypes.h"
#include <algorithm>
#include <cassert>

namespace leetcode {
	class Solution {
	public:
		int maxDepth(TreeNode* root) {
			if (root == nullptr)
				return 0;

			int depth = 1;
			return depth + std::max(maxDepth(root->left), maxDepth(root->right));
		}

		void maxDepthTest()
		{
			TreeNode* root = new TreeNode(1);
			root->left = new TreeNode(2);
			root->right = new TreeNode(3);
			root->left->left = new TreeNode(4);
			root->left->left->left = new TreeNode(5);

			assert(4 == maxDepth(root));
		}
	};
}
