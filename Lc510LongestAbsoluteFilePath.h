#pragma once

#include <string>
#include <stack>
#include <vector>
#include <cassert>

#define PRINT_MAXPATH

#ifdef PRINT_MAXPATH
#include <iostream>
#endif

using namespace std;

namespace leetcode
{
// the question url:https://leetcode.com/problems/longest-absolute-file-path/
class Solution
{
  public:
    struct Node
    {
        string value = "";
        unsigned int level = 0;
        bool isFile = false;
    };

  public:
    int lengthLongestPath(string input)
    {
        unsigned int pos = 0;
        size_t len = 0;
        vector<string> path;
        size_t maxLen = 0;
        vector<string> maxPath;
        stack<Node *> stack;
        while (pos < input.length())
        {
            Solution::Node *n = new Solution::Node();
            readNode(n, input, pos);
            while (!stack.empty())
            {
                auto ntop = stack.top();
                if (ntop->level >= n->level)
                {
                    stack.pop();
                    path.pop_back();
                    len -= ntop->value.length();
                }
                else
                {
                    break;
                }
            }

            if (n->isFile)
            {
                if (len + n->value.length() > maxLen)
                {
                    maxPath = vector<string>(path);
                    maxPath.push_back(n->value);
                    maxLen = len + n->value.length();
                }
            }
            else
            { // is a dir
                stack.push(n);
                path.push_back(n->value);
                len += n->value.length();
            }
        }

#ifdef PRINT_MAXPATH
        for (size_t i = 0; i < maxPath.size(); i++)
        {
            cout << maxPath[i];
            if (i != maxPath.size() - 1)
                cout << "/";
        }
#endif

        if (maxPath.size() > 1)
            maxLen += maxPath.size() - 1; // consider the path seperator chars "/"

        return (int)maxLen;
    }

    void readNode(Node *n, string &input, unsigned int &pos)
    {
        while (pos < input.length())
        {
            auto c = input[pos];
            pos++;
            if (c == '\n')
            { //end of this node
                break;
            }
            else if (c == '\t')
            {
                n->level++;
            }
            else
            {
                n->value.push_back(c);
                if (c == '.')
                    n->isFile = true;
            }
        }
    }

    void lengthLongestPathTest()
    {
        string input = "dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext";
        assert(lengthLongestPath(input) == 32);

        string input2 = "A";
        assert(lengthLongestPath(input2) == 0);

        string input3 = "a\n\taa\n\t\taaa\n\t\t\tfile1234567890123.txt\naaaaaaaaaaaaaaaaaaaaa\n\tsth.png";
        auto len = lengthLongestPath(input3);
        assert(len == 30);

        string input_name_with_spaces = "dir\n    file.txt";
        assert(lengthLongestPath(input_name_with_spaces) == 12);

        string input_name_with_spaces_2 = "file name with  space.txt";
        assert(lengthLongestPath(input_name_with_spaces_2) == 25);
    }
};
}
