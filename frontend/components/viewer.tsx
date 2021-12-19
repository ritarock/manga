import {BookData} from '../interfaces'

export default function Viewer({booksData}: {booksData: BookData[]}) {
  return (
    <>
      <div className='box-row'>
        {booksData.map(({Isbn, Title, Cover}) => (
          <div key={Isbn}>
            {/* eslint-disable-next-line @next/next/no-img-element */}
            <img src={Cover} alt={Title} width="" height=""></img>
          </div>
        ))}
      </div>
      <style jsx>{`
        .box-row {
          display: flex;
          display: -webkit-box;
          display: -ms-flexbox;
          flex-wrap: wrap
        }
      `}
      </style>
    </>
  )
}
