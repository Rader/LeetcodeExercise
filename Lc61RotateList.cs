using System;


namespace LeetcodeExercise
{
    public class Lc61RotateList
    {
        public ListNode RotateRight(ListNode head,int k)
        {
            if (head == null || k == 0)
                return head;

            var leftHead = head;
            var leftTail = leftHead;

            var rightHead = head;
            var rightTail = rightHead;

            int len = 1;
            while (len<k)
            {
                rightTail = rightTail.next;
                if(rightTail==null)
                {
                    k %= len;

                    len = 1;
                    rightTail = rightHead;
                }
                else
                {
                    len++;
                }
            }
            if (k == 0) // no rotate is needed
                return head;

            // begine rotate
            bool hasRotated = false;
            while (rightTail.next!=null)
            {
                if(!hasRotated)
                    hasRotated = true;

                leftTail = rightHead;

                rightHead = rightHead.next;
                rightTail = rightTail.next;
            }

            if (hasRotated)
            {
                leftTail.next = null;
                rightTail.next = leftHead;
            }

            return rightHead;
        }
    }


    public class ListNode
    {
      public int val;
      public ListNode next;
      public ListNode(int x) { val = x; }
  }
}
