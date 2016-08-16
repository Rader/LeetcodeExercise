using NUnit.Framework;

namespace LeetcodeExercise.Tests
{
    [TestFixture]
    public class Lc61RotateListTest
    {
        [Test]
        public void RotateRight_Empty_List()
        {
            Lc61RotateList obj = new Lc61RotateList();
            ListNode head = null;
            Assert.AreEqual(null, obj.RotateRight(head, 1));
        }

        [Test]
        public void RotateRight_One()
        {
            Lc61RotateList obj = new Lc61RotateList();
            ListNode head = BuildList(new int[] { 1, 2, 3 });
            int k = 1;
            ListNode expectedHead = BuildList(new int[] { 3, 1, 2 });
            ListNode actualHead = obj.RotateRight(head, k);
            Assert.IsTrue(DeepEqual(expectedHead, actualHead));

        }

        [Test]
        public void RotateRight_Two()
        {
            Lc61RotateList obj = new Lc61RotateList();
            ListNode head = BuildList(new int[] { 1, 2, 3 });
            int k = 2;
            ListNode expectedHead = BuildList(new int[] { 2, 3, 1 });
            ListNode actualHead = obj.RotateRight(head, k);
            Assert.IsTrue(DeepEqual(expectedHead, actualHead));

        }

        [Test]
        public void RotateRight_Three()
        {
            Lc61RotateList obj = new Lc61RotateList();
            ListNode head = BuildList(new int[] { 1, 2, 3 });
            int k = 3;
            ListNode expectedHead = BuildList(new int[] { 1, 2, 3 });
            ListNode actualHead = obj.RotateRight(head, k);
            Assert.IsTrue(DeepEqual(expectedHead, actualHead));

        }

        [Test]
        public void RotateRight_Four()
        {
            Lc61RotateList obj = new Lc61RotateList();
            ListNode head = BuildList(new int[] { 1, 2, 3 });
            int k = 4;
            ListNode expectedHead = BuildList(new int[] { 3, 1, 2 });
            ListNode actualHead = obj.RotateRight(head, k);
            Assert.IsTrue(DeepEqual(expectedHead, actualHead));

        }

        [Test]
        public void RotateRight_Ten()
        {
            Lc61RotateList obj = new Lc61RotateList();
            ListNode head = BuildList(new int[] { 1, 2, 3, 4, 5 });
            int k = 10;
            ListNode expectedHead = BuildList(new int[] { 1,2,3,4,5 });
            ListNode actualHead = obj.RotateRight(head, k);
            Assert.IsTrue(DeepEqual(expectedHead, actualHead));
        }

        [Test]
        public void RotateRight_More()
        {
            Lc61RotateList obj = new Lc61RotateList();
            ListNode head = BuildList(new int[] { 1, 2, 3 });
            int k=2000000000;
            ListNode expectedHead = BuildList(new int[] { 2, 3, 1 });
            ListNode actualHead = obj.RotateRight(head, k);
            Assert.IsTrue(DeepEqual(expectedHead, actualHead));
        }

        public static bool DeepEqual(ListNode one,ListNode two)
        {
            if (one == two)
                return true;

            if (one == null || two == null)
                return false;

            if (one.val != two.val)
                return false;

            return DeepEqual(one.next, two.next);
        }

        public static ListNode BuildList(int[] values)
        {
            if (values == null || values.Length == 0)
                return null;

            ListNode head = new ListNode(values[0]);
            var tail = head;
            for (int i = 1; i < values.Length; i++)
            {
                tail.next = new ListNode(values[i]);
                tail = tail.next;
            }

            return head;
        }
    }
}
