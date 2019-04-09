import pandas as pd

inp = '''
>Rosalind_1
ATCCAGCT
>Rosalind_2
GGGCAACT
>Rosalind_3
ATGGATCT
>Rosalind_4
AAGCAACC
>Rosalind_5
TTGGAACT
>Rosalind_6
ATGCCATT
>Rosalind_7
ATGGCACT
'''.strip()

def loadData(fn):

    with open(fn, 'r') as f:
        vals = [val.strip() for val in f]

    df = pd.DataFrame()
    s = []

    for val in vals:
        if val[0] != '>':
            s.extend(list(val))
        else:
            s = { i:v for i,v in enumerate(s) }
            df = df.append(s, ignore_index=True)
            s = []
    s = { i:v for i,v in enumerate(s) }
    df = df.append(s, ignore_index=True)
    return df.iloc[1:,:]

def getColCounts(df, c):
    ls = []
    for column in df:
        ls.append(str(df[column][df[column] == c].count()))
    return ls

df = loadData('rosalind_cons.txt')
na = 'ACGT'

most = ''.join(list(df.mode().iloc[0,:]))
print(most)
for n in na:
    print(n + ': ', end='')
    print(' '.join(getColCounts(df,n)))

getColCounts(df, 'A')
