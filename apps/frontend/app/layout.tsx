import '../styles/globals.css';

export const metadata = {
    title: 'Adapticc',
    description: 'Built by FollowThePattern',
}

export default function RootLayout({
    children,
}: {
    children: React.ReactNode
}) {
    return (
        <html lang="en" className='h-full bg-white'>
            <body className="h-full">
                {children}
            </body>
        </html>
    )
}
