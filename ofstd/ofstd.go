package ofstd

/** check whether the addition of two 32-bit integers yields in an overflow
 *  @param summand1 first integer value to be added
 *  @param summand2 second integer value to be added
 *  @return OFTrue if an overflow occurred during the addition, OFFalse otherwise
 */
func Check32BitAddOverflow(summand1 uint32, summand2 uint32) bool {
	return (0xffffffff-summand1 < summand2)
}
