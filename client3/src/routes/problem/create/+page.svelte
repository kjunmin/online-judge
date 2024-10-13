<script lang="ts">
	import { Editor } from '@tadashi/svelte-editor-quill';
	import { Button, Input, Label } from 'flowbite-svelte';

	const options = {
		theme: 'snow'
	};
	let data = '<p>Insert description...</p>';

	async function handleSubmit(event: Event) {
		const formEl = event.target as HTMLFormElement;
		const formData = new FormData(formEl);

		const formObj = Object.fromEntries(formData.entries());
		console.log({ formObj });
	}

	const FormFields = {
		Title: 'title',
		Description: 'description'
	} as const;
</script>

<svelte:head>
	<link rel="stylesheet" href="https://unpkg.com/quill@2.0.2/dist/quill.snow.css" crossorigin />
</svelte:head>

<h1>Create problem</h1>

<form on:submit|preventDefault={handleSubmit}>
	<Label for={FormFields.Title}>Title:</Label>
	<Input type="text" required id={FormFields.Title} name={FormFields.Title} />
	<Label for={FormFields.Description}>Description:</Label>
	<input
		type="text"
		required
		hidden
		id={FormFields.Description}
		name={FormFields.Description}
		bind:value={data}
	/>
	<Editor
		{data}
		{options}
		on:text-change={(event) => {
			data = event?.detail?.html ?? '';
		}}
	/>
	<Button type="submit">Submit</Button>
</form>
