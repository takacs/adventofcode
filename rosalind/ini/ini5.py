with open('rosalind_ini5.txt', 'r') as f:
    for i, line in enumerate(f, start=1):
        if i%2 == 0:
            print(line.strip())
