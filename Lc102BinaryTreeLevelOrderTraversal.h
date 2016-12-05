#pragma once

#include "LcTypes.h"
#include <vector>
using namespace std;

namespace leetcode {
	class Solution {
	public:
		/* Apply breadth-first traversal.*/
		vector<vector<int>> levelOrder(TreeNode* root) {			
			vector<vector<int>> vlevels;
			if (root == nullptr)
				return vlevels;

			vector<TreeNode*> parents;			
			parents.push_back(root);
			while (!parents.empty())
			{
				vector<TreeNode*> children;
				vector<int> vlevel;				
				for (auto pnode : parents)
				{
					vlevel.push_back(pnode->val);					
					if (pnode->left != nullptr) children.push_back(pnode->left);
					if (pnode->right != nullptr) children.push_back(pnode->right);
				}

				vlevels.push_back(vlevel);
				parents = children;
			}	

			return vlevels;
		}
	};
}
