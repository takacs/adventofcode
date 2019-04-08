def sliceword(s,a,b,c,d):
    return ' '.join([s[a:b+1], s[c:d+1]])

s = 'DwVmicySJ91AkWeBIDixUt7CfPOF4mkg5uE7AcheronMV3D4GFJnKSdThgkJf7zHrDkrPpnlu13snnuptaAAXKFEuaRENjFRRBytTzbsUcAXytLPJtust1UXb7G4ra9HKeKo7W5xe2Jc20RRBUybyavToLO7tVdFkK6HqD3Ph2xU30FZ6uoSEXi.'
a,b,c,d = 36,42,77,81

print(sliceword(s,a,b,c,d))
print(sliceword('HumptyDumptysatonawallHumptyDumptyhadagreatfallAlltheKingshorsesandalltheKingsmenCouldntputHumptyDumptyinhisplaceagain.', 22, 27, 97, 102))
