import org.junit.Assert;
import org.junit.Test;

public class Lc540SingleElementInSortedArrayTest {

	@Test
	public void testFindSingle() {
		int[] vals = new int[] { 1, 1, 2, 2, 3, 3, 4, 4, 5, 6, 6 };
		Lc540SingleElementInSortedArray obj = new Lc540SingleElementInSortedArray();
		int single = obj.findSingle(vals);
		Assert.assertEquals(5, single);

		vals = new int[] { 1, 1, 2, 3, 3, 4, 4, 8, 8 };
		single = obj.findSingle(vals);
		Assert.assertEquals(2, single);

		vals = new int[] { 3, 3, 7, 7, 10, 11, 11 };
		single = obj.findSingle(vals);
		Assert.assertEquals(10, single);

		vals = new int[] { 3, 7, 7, 10, 10, 11, 11 };
		single = obj.findSingle(vals);
		Assert.assertEquals(3, single);

		vals = new int[] { 3, 3, 7, 10, 10, 11, 11 };
		single = obj.findSingle(vals);
		Assert.assertEquals(7, single);

		vals = new int[] { 3, 3, 7, 7, 10, 10, 11 };
		single = obj.findSingle(vals);
		Assert.assertEquals(11, single);
	}

}
