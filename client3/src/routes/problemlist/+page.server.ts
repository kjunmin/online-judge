export async function load() {
	const response = await fetch('/api/v1/problemslist');
	const problemList = await response.json();

	return {
		problemList
	};
}
