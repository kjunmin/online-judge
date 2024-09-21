import { Button, TagsInput, TextInput } from "@mantine/core";
import styles from './CreateProblem.module.css'
import { useForm } from "@mantine/form";
import { Problem } from "../types";
import axios from "axios";

function useCreateProblem() {
    type CreateProblem = Omit<Problem, 'problemID'>

    function createProblem(problem: CreateProblem) {
        axios.post('/api/v1/problem/create', problem)
    }

    return { createProblem };
}

export function CreateProblem() {

    const { createProblem } = useCreateProblem();

    const form = useForm({
        mode: 'uncontrolled',
        initialValues: {
            tags: [],
            title: '',
            description: '',
        },
    });

    function handleSubmit(problem: typeof form.values) {
        createProblem(problem)
    }

    return (
        <div className={styles.container}>
            <form onSubmit={form.onSubmit(handleSubmit)}>
                <TextInput
                    required
                    key={form.key('title')}
                    label="Title"
                    placeholder="Remove Nth Node from end of list"
                    {...form.getInputProps('title')}
                />
                <TextInput
                    required
                    key={form.key('description')}
                    label="Description"
                    {...form.getInputProps('description')}
                />
                <TagsInput
                    key={form.key('tags')}
                    label="Tags"
                    {...form.getInputProps('tags')}
                />
                <Button type="submit">Submit</Button>
            </form>
        </div>
    )
}