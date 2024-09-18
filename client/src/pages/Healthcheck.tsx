import { useQuery } from '@tanstack/react-query'
import axios from 'axios'
import React from 'react'
import { useEffect } from 'react'

type HealthCheck = {
    status: 'ok' | 'overloaded'
}

function useHealthCheck() {
    return useQuery({
        queryKey: ['retrieveConn'],
        queryFn: async () => {
            const response = await axios.get<never, HealthCheck>('/api/v1/healthcheck')
            return response
        }
    })
}

export function HealthCheck() {
    const { refetch, data } = useHealthCheck()

    useEffect(() => {
        const interval = setInterval(refetch, 5000);

        return () => {
            clearInterval(interval)
        }
    }, [refetch])


    return (
        <div>
            <h1>{`Status: ${data?.status}`}</h1>
        </div>
    )
}