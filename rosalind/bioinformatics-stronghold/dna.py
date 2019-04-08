from collections import Counter

def stringcounter(s):
    return Counter(s)

occs = stringcounter('TCTTTTGTATTCGCTATTAACTTTACATTTCCCAAATACAGTTGAGAGGTAACCGACTTCACCAGTCGTTCCCCTTTTTCTTAATTGCTCCGAAAGACATTTACCCGTATAGGATATAAAAAGCGCGAATCCTAGGTTTTGGTGGTGGCGAAAGAAATATGGCAGACCGACGCTTCGCGAGTAAGTGTTACCCCGGGGGGCGCTAGTCGGATCGGATGACAGTGTTCGAACAGACCACGAGGGAAAAGAGCGAAGACTTACATCACGCGAAAAGTTTAAGAGCATCAATGTCTTTCTCACGCAATCAAAGTTCCTCGATGTGTCAGCGCACGCCCAAAAGCGTTCGGGTGCCTGACTGGGCCGTTGCATATATGCGCACTAATAATGACCGTTAAGCCAAGTGCGAATAATTGACGTTAAGCATAGCAAAGTTTCTGTAGAATGTTATCTAAGATAGTTATAAATACGCGACCTGCTGACTCAGTTCTTCTTTTCGACGGATTGTACGAGTCGCACGAGGCTGCCCTGTTTACAAAGTTTCTGTGCTGTACCGAAACAATCGCTGGAACCATCGAGGAGGTTATCTAACTAGCATTCATCGTCGAATGAGTCACGATAATGGACCTTCATAGATCTCTTGGCGGTACTTGCCCTTCCACATCAATAGCATGCAGAAAAACGTAATGCTAATTTCAGACGTTTTTGCCCGGGCTTTTAAGGTCTTACGGTTTGCCACTGTCTAAGGTGAGATCATATCCGAGTGTGGGCGCGAATTCAGGAGGAGTTGCAGTTATCATAGAACAGAACCTTTACTGGTGACCATCGCAATTATGTTTGACTGCTCTTATCAGCTCATCGCATTAGTTTTAGCACCTAATTTCCTTCTCTAATTCATATAGT'.strip())

for k in sorted(occs.keys()):
    print(occs[k], end=' ')