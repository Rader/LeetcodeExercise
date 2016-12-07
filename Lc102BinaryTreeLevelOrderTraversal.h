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
		
		vector<vector<int>> levelOrder_SingleQueue(TreeNode* root) {
			vector<vector<int>> vlevels;
			if (root == nullptr)
				return vlevels;

			queue<TreeNode*> queue;
			queue.push(root);
			while (!queue.empty())
			{
				int parentsCount = queue.size();
				vector<int> vlevel;
				while (parentsCount>0)
				{
					TreeNode* parent = queue.front();
					vlevel.push_back(parent->val);
					if (parent->left != nullptr) queue.push(parent->left); 
					if (parent->right != nullptr) queue.push(parent->right);

					queue.pop();
					parentsCount--;
				}
				vlevels.push_back(vlevel);
			}	

			return vlevels;
		}
		
		void levelOrderTest() {
			TreeNode* root = new TreeNode(1);
			root->left = new TreeNode(2);
			root->right = new TreeNode(3);
			root->left->left = new TreeNode(4);
			root->left->left->left = new TreeNode(5);

			vector<vector<int>> expected;
			expected.push_back(vector<int>{1});
			expected.push_back(vector<int>{2, 3});
			expected.push_back(vector<int>{4});
			expected.push_back(vector<int>{5});

			auto result = levelOrder(root);
			assert(expected == result);

			auto result2 = levelOrder_SingleQueue(root);
			assert(expected == result2);
		}
	};
}
