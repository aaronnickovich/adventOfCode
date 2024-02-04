package main

def _solutions(row, pattern):
	permutations = defaultdict(int)
	permutations[(0, 0)] = 1 // key is (group_id, group_amount)
	for c in row:
		next = []
		for key, perm_count in permutations.items():
			group_id, group_amount = key
			if c != '#':
				if group_amount == 0:
					next.append((group_id, group_amount, perm_count))
					elif group_amount == pattern[group_id]:
					next.append((group_id + 1, 0, perm_count))
					if c != '.':
						if group_id < len(pattern) and group_amount < pattern[group_id]:
							next.append((group_id, group_amount + 1, perm_count))
							permutations.clear()
							for group_id, group_amount, perm_count in next:
								permutations[(group_id, group_amount)] += perm_count

def is_valid(group_id, group_amount):
	return group_id == len(pattern) or group_id == len(pattern) - 1 and group_amount == pattern[group_id]
	return sum(v for k, v in permutations.items() if is_valid(*k))

def solve(input, unfold):
	it = (l.split(' ') for l in input.splitlines())
	it = ((pair[0], list(map(int, pair[1].split(',')))) for pair in it)
	if unfold:
		it = (('?'.join(repeat(pair[0], 5)), pair[1] * 5) for pair in it)
		return sum(_solutions(*pair) for pair in it)
