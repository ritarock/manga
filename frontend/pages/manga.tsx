import {GetStaticProps} from 'next'
import Viewer from '../components/viewer'
import {BookData} from '../interfaces'


interface Book {
  code: number
  data: BookData[]
}

export default function Manga({books}: {books: Book}) {
  return (
    <>
      <Viewer booksData={books.data} />
    </>
  )
}

export const getStaticProps: GetStaticProps = async () => {
  const res = await fetch("http://0.0.0.0:8080/backend/manga")
  const books = await res.json()

  return {
    props: {
      books,
    }
  }
}
