<template>
    <a-result status="404" subtitle="Whoops, that page is gone(404).">
        <template #extra>
            <a-space>
                <a-button type="primary" @back="$router.go(-1)">Back</a-button>
            </a-space>
        </template>
    </a-result>
</template>
