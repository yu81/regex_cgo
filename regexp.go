package cregexp

/*
# include <regex.h>
int regmatch(char* target, char* pattern, int cflags) {
	regex_t reg;
	regmatch_t match;
	regcomp(&reg, pattern, cflags);
	int execRet = regexec(&reg, target, 1, &match, 0);
	regfree(&reg);
	return execRet;
}
*/
import "C"

const (
	// regmatch options.
	REG_EXTENDED = 1
	REG_ICASE    = 2
	REG_NOSUB    = 4
	REG_NEWLINE  = 8
)

const (
	// regexec return values.
	REG_OKAY = iota
	REG_NOMATCH
	REG_BADPAT
	REG_ECOLLAT
	REG_ECTYPE
	REG_EESCAPE
	REG_ESUBREG
	REG_EBRACK
	REG_EPAREN
	REG_EBRACE
	REG_BADBR
	REG_ERANGE
	REG_ESPACE
	REG_BADRPT
	REG_EMPTY
	REG_ASSERT
	REG_INVARG
	REG_ATOI = 255  /* convert name to number (!) */
	REG_ITOA = 0400 /* convert number to name (!) */
)

func RegMatch(target, pattern string, cFlags int) bool {
	result := C.regmatch(C.CString(target), C.CString(pattern), C.int(cFlags))
	switch result {
	case REG_OKAY:
		return true
	default:
		return false
	}
}
