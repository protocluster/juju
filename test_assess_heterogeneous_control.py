__metaclass__ = type

from argparse import Namespace
import os
from unittest import TestCase

from mock import patch

from assess_heterogeneous_control import (
    dumping_env,
    get_clients,
    parse_args,
    upload_heterogeneous,
    )
from jujupy import (
    EnvJujuClient,
    SimpleEnvironment,
    _temp_env,
    )


class TestDumping_env(TestCase):

    def test_dumping_env_exception(self):
        client = EnvJujuClient(SimpleEnvironment('env'), '5', '4')
        with patch.object(client, 'destroy_environment') as de_mock:
            with patch('subprocess.check_call') as cc_mock:
                with patch('deploy_stack.copy_remote_logs') as crl_mock:
                    with self.assertRaises(ValueError):
                        with dumping_env(client, 'foo', 'bar'):
                            raise ValueError
        crl_mock.assert_called_once_with('foo', 'bar')
        cc_mock.assert_called_once_with(['gzip', '-f'])
        de_mock.assert_called_once_with()

    def test_dumping_env_success(self):
        client = EnvJujuClient(SimpleEnvironment('env'), '5', '4')
        with patch.object(client, 'destroy_environment') as de_mock:
            with patch('subprocess.check_call') as cc_mock:
                with patch('deploy_stack.copy_remote_logs') as crl_mock:
                    with dumping_env(client, 'foo', 'bar'):
                        pass
        crl_mock.assert_called_once_with('foo', 'bar')
        cc_mock.assert_called_once_with(['gzip', '-f'])
        self.assertEqual(de_mock.call_count, 0)


class TestParseArgs(TestCase):

    def test_parse_args(self):
        args = parse_args(['a', 'b', 'c', 'd', 'e'])
        self.assertEqual(args, Namespace(
            initial='a', other='b', base_environment='c',
            environment_name='d', log_dir='e', debug=False,
            upload_tools=False, agent_url=None,
            user=os.environ.get('JENKINS_USER'),
            password=os.environ.get('JENKINS_PASSWORD')))

    def test_parse_args_agent_url(self):
        args = parse_args(['a', 'b', 'c', 'd', 'e', '--agent-url', 'foo',
                           '--user', 'my name', '--password', 'fake pass'])
        self.assertEqual(args.agent_url, 'foo')
        self.assertEqual(args.user, 'my name')
        self.assertEqual(args.password, 'fake pass')


class TestGetClients(TestCase):

    def test_get_clients(self):
        boo = {
            ('foo', '--version'): '1.18.73',
            ('bar', '--version'): '1.18.74',
            ('juju', '--version'): '1.18.75',
            ('which', 'juju'): '/usr/bun/juju'
            }
        with _temp_env({'environments': {'baz': {}}}):
            with patch('subprocess.check_output', lambda x: boo[x]):
                initial, other, released = get_clients('foo', 'bar', 'baz',
                                                       'qux', True, 'quxx')
        self.assertEqual(initial.env, other.env)
        self.assertEqual(initial.env, released.env)
        self.assertEqual(initial.env.config['tools-metadata-url'], 'quxx')
        self.assertEqual(initial.full_path, os.path.abspath('foo'))
        self.assertEqual(other.full_path, os.path.abspath('bar'))
        self.assertEqual(released.full_path, '/usr/bun/juju')

    def test_get_clients_no_agent(self):
        with _temp_env({'environments': {'baz': {}}}):
            with patch('subprocess.check_output', return_value='1.18.73'):
                initial, other, released = get_clients('foo', 'bar', 'baz',
                                                       'qux', True, None)
        self.assertTrue('tools-metadata-url' not in initial.env.config)


class TestUploadHeterogeneous(TestCase):

    def test_upload_heterogeneous(self):
        args = Namespace(user='foo', password='bar')
        env_ctx = patch(
            'upload_hetero_control.HUploader.upload_by_env_build_number')
        with patch('upload_hetero_control.get_s3_access',
                   return_value=('name', 'pass')) as a_mock:
            with patch('upload_hetero_control.S3Connection') as c_mock:
                with env_ctx as u_mock:
                    upload_heterogeneous(args)
        a_mock.assert_called_once_with()
        c_mock.assert_called_once_with('name', 'pass')
        u_mock.assert_called_once_with()
