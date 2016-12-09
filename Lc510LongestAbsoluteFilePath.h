#pragma once

#include <string>
#include <stack>
using namespace std;

namespace leetcode {
	class Solution {
	public:
		struct Node
		{
			string value="";
			unsigned int level = 0;
		};

	public:
		int lengthLongestPath(string input) {
			unsigned int pos = 0;
			unsigned int curr_level = 0;
			stack<Node*> stack;
			while (pos < input.length())
			{
				Solution::Node* n = new Solution::Node();
				readNode(n, input, pos);
				if (stack.empty()) {
					stack.push(n);
				}
				else
				{

				}
			}
		}

		void readNode(Node* n, string& input, unsigned int& pos) {
			while (pos < input.length())
			{
				auto c = input[pos];
				pos++;
				if (c == '\n') { //end of this node
					break;
				}
				else if (c == '\t'){
					n->level++;
				}
				else
				{
					n->value.push_back(c);
				}				
			}
		}


		void lengthLongestPathTest() {
			string input = "dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext";
		}

	
	};
}
