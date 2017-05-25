import React from 'react'
import PropTypes from 'prop-types'
import {injectIntl, intlShape} from 'react-intl'

const Widget = ({intl}) => (
  <footer>
    <p className="float-right"><a href="/">Back to top</a></p>
    <p>
      &copy; {intl.formatMessage({id: "site.copyright"})}
      &middot; <a href="/">Privacy</a>
      &middot; <a href="/">Terms</a>
    </p>
  </footer>
)

Widget.propTypes = {
    intl: intlShape.isRequired
}

export default injectIntl(Widget)
