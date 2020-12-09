import sys
import re

def parse(inp):
    allpass = []
    currpass = ''
    for line in inp:
        if line == '':
            keyvals = currpass.split()
            passdict = { kv.split(':')[0]:kv.split(':')[1] for kv in keyvals }
            allpass.append(passdict)
            currpass = ''
        else:
            currpass += ' ' + line
    return allpass

def countValids(ps):
    fields = ['byr', 'iyr', 'eyr', 'hgt', 'hcl', 'ecl', 'pid']
    count = 0
    for p in ps:
        if all([f in p for f in fields]):
            count += 1
    return count

def countSuperValids(ps):
    fields = ['byr', 'iyr', 'eyr', 'hgt', 'hcl', 'ecl', 'pid']
    count = 0
    for p in ps:
        if all([f in p for f in fields]) and superValid(p):
            count += 1
    return count
    
def superValid(p):
    pattern = re.compile("^(#[0-9a-f]{6}$)")
    passpattern = re.compile("^([0-9]{9}$)")
    valchecks = []
    valchecks.append(1920 <=  int(p['byr']) <= 2002)
    valchecks.append(2010 <=  int(p['iyr']) <= 2020)
    valchecks.append(2020 <=  int(p['eyr']) <= 2030)
    valchecks.append(
        (p['hgt'][-2:] == 'cm' and 150 <= int(p['hgt'][:-2]) <= 193) or
        (p['hgt'][-2:] == 'in' and 59 <= int(p['hgt'][:-2]) <= 76)
        )
    valchecks.append(bool(pattern.match(p['hcl'])))
    valchecks.append(p['ecl'] in ['amb', 'blu', 'brn', 'gry', 'grn', 'hzl', 'oth'])
    valchecks.append(bool(passpattern.match(p['pid'])))
    return all(valchecks)
    
inp = open(sys.argv[1]).read().strip().split('\n')
inp.append('')
passports = parse(inp)
print(countValids(passports))
print(countSuperValids(passports))
