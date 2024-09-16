import { useQuery } from '@tanstack/react-query'

function useTests() {
    return useQuery({
        queryKey: ['retrieveConn'],
        queryFn: async (): Promise<unknown> => {
            const response = await fetch('http://localhost:8080/api/v1/healthcheck')
            return await response.json()
        }
    })
}

export function Test() {
    const { refetch } = useTests()

    return <div>
        <h1>Test</h1>
        <button onClick={() => refetch()}>Test connectivity</button>
    </div>
}