import React from 'react'
import { Button, Form, Input } from 'reactstrap'
import {injectIntl, intlShape} from 'react-intl'

const Widget = ({intl}) => (
  <Form className="mt-2 mt-md-0" inline>
    <Input className="mr-sm-2" type="text" id="keyword" placeholder={intl.formatMessage({id: 'hints.search'})} />
    <Button outline color="success" className="my-2 my-sm-0">
      {intl.formatMessage({id: 'buttons.search'})}
    </Button>
  </Form>
)

Widget.propTypes = {
    intl: intlShape.isRequired
}

export default injectIntl(Widget)
